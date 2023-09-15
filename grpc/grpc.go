package grpc

import (
	"context"
	"net"
	"strings"
	"sync"
	"time"

	"example.com/go-inmem-loki/api"
	"example.com/go-inmem-loki/constant"
	"example.com/go-inmem-loki/inmemory"
	"example.com/go-inmem-loki/monitoring"
	pb "example.com/go-inmem-loki/proto"
	"example.com/go-inmem-loki/types"
	"example.com/go-inmem-loki/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	m      *inmemory.InMemoryStore
	locker *util.LockManager
	pb.UnimplementedFeatureStoreServer
}

func (s *server) HeartBeat(ctx context.Context, req *pb.PingReq) (*pb.PongResp, error) {
	return &pb.PongResp{
		Pong: "pong",
	}, nil
}

func (s *server) GetSkuFeatures(ctx context.Context, req *pb.GetSkuFeaturesReq) (*pb.GetSkuFeaturesResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()
	skuId := req.GetSkuId()
	catalogueName := req.GetCatalogueName()

	feat := s.m.GetSkuFeatures(skuId, catalogueName, traceId)

	elapsed := time.Since(startTime).Microseconds()
	monitoring.RecordMetrics("GetSkuFeatures", elapsed)
	log.Info().Str("Component", "GetSkuFeatures").Str("TraceId:", traceId).Msgf("GetSkuFeatures req %v took %v microseconds", req.GetSkuId(), elapsed)

	return &pb.GetSkuFeaturesResp{Features: feat}, nil
}

func (s *server) GetSkuFeaturesForAllCatalogues(ctx context.Context, req *pb.GetSkuFeaturesForAllCataloguesReq) (*pb.GetSkuFeaturesForAllCataloguesResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()
	skuId := req.GetSkuId()

	m := s.m.GetSkuFeaturesForAllCatalogues(skuId, traceId)

	elapsed := time.Since(startTime).Microseconds()
	monitoring.RecordMetrics("GetSkuFeaturesForAllCatalogues", elapsed)
	log.Info().Str("Component", "GetSkuFeaturesForAllCatalogues").Str("TraceId:", traceId).Msgf("GetSkuFeaturesForAllCatalogues req %v took %v microseconds", req.GetSkuId(), elapsed)

	return &pb.GetSkuFeaturesForAllCataloguesResp{FeaturesMap: m}, nil
}

func (s *server) Lock(ctx context.Context, req *pb.LockReq) (*pb.LockResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()
	key := req.GetKey()
	ttl := req.GetTtl()
	owner := req.GetOwner()
	defer func() {
		elapsed := time.Since(startTime).Microseconds()
		monitoring.RecordMetrics("Lock", elapsed)
		log.Info().Str("Component", "Lock").Str("TraceId:", traceId).Msgf("Lock req %v took %v microseconds", req.GetKey(), elapsed)
	}()

	ok := s.locker.AcquireLock(key, owner, time.Until(time.Now().Add(time.Duration(ttl)*time.Second)))
	if !ok {
		log.Error().Str("Component", "Lock").Str("TraceId:", traceId).Msgf("failed to acquire lock: %v", key)
		return nil, status.Errorf(codes.Internal, "failed to acquire lock: %v", key)
	}

	return &pb.LockResp{}, nil
}

func (s *server) Release(ctx context.Context, req *pb.ReleaseReq) (*pb.ReleaseResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()
	key := req.GetKey()
	owner := req.GetOwner()
	s.locker.ReleaseLock(key, owner)

	elapsed := time.Since(startTime).Microseconds()
	monitoring.RecordMetrics("Release", elapsed)
	log.Info().Str("Component", "Release").Str("TraceId:", traceId).Msgf("Release req %v took %v microseconds", req.GetKey(), elapsed)

	return &pb.ReleaseResp{}, nil
}

func (s *server) UpsertSkuFeatures(ctx context.Context, req *pb.UpsertSkuFeaturesReq) (*pb.UpsertSkuFeaturesResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()

	defer func() {
		elapsed := time.Since(startTime).Microseconds()
		monitoring.RecordMetrics("UpsertSkuFeatures", elapsed)
		log.Info().Str("Component", "UpsertSkuFeatures").Str("TraceId:", traceId).Msgf("UpsertSkuFeatures req %v took %v microseconds", req.GetSkuId(), elapsed)
	}()

	skuId := req.GetSkuId()
	catalogueName := req.GetCatalogueName()
	features := req.GetFeatures()
	var ttl *int
	if req.Ttl != nil {
		v := int(*req.Ttl)
		ttl = &v
	}

	payload := types.UpdateSkuPayload{
		SkuId:         skuId,
		CatalogueName: catalogueName,
		TTLSeconds:    ttl,
		Features:      features,
	}

	err := s.m.UpsertSkuFeatures(payload)

	if err != nil {
		log.Error().Str("Component", "UpsertSkuFeatures").Str("TraceId:", traceId).Msgf("failed to upsert sku features: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to upsert sku features: %v", err)
	}

	return &pb.UpsertSkuFeaturesResp{}, nil
}

func (s *server) OrderExists(ctx context.Context, req *pb.OrderExistsReq) (*pb.OrderExistsResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()

	defer func() {
		elapsed := time.Since(startTime).Microseconds()
		monitoring.RecordMetrics("OrderExists", elapsed)
		log.Info().Str("Component", "OrderExists").Str("TraceId:", traceId).Msgf("OrderExists req %v took %v microseconds", req.GetSkuId(), elapsed)
	}()

	skuId := req.GetSkuId()
	userId := req.GetUserId()

	val := s.m.OrderExists(skuId, userId)

	return &pb.OrderExistsResp{Value: val}, nil
}

func (s *server) MarkOrder(ctx context.Context, req *pb.MarkOrderReq) (*pb.MarkOrderResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()

	defer func() {
		elapsed := time.Since(startTime).Microseconds()
		monitoring.RecordMetrics("MarkOrder", elapsed)
		log.Info().Str("Component", "MarkOrder").Str("TraceId:", traceId).Msgf("MarkOrder req %v took %v microseconds", req.GetSkuId(), elapsed)
	}()

	skuId := req.GetSkuId()
	userId := req.GetUserId()
	expiresAt, err := time.Parse(time.RFC3339, req.GetExpiresAt())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse expiresAt: %v", err)
	}

	err = s.m.MarkOrder(skuId, userId, expiresAt)

	return &pb.MarkOrderResp{}, err
}

func (s *server) ImpressionExists(ctx context.Context, req *pb.ImpressionExistsReq) (*pb.ImpressionExistsResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()

	defer func() {
		elapsed := time.Since(startTime).Microseconds()
		monitoring.RecordMetrics("ImpressionExists", elapsed)
		log.Info().Str("Component", "ImpressionExists").Str("TraceId:", traceId).Msgf("ImpressionExists req %v took %v microseconds", req.GetSkuId(), elapsed)
	}()

	skuId := req.GetSkuId()
	userId := req.GetUserId()

	val := s.m.GetImpressionExists(skuId, userId)

	return &pb.ImpressionExistsResp{Value: val}, nil
}

func (s *server) MarkImpression(ctx context.Context, req *pb.MarkImpressionReq) (*pb.MarkImpressionResp, error) {
	startTime := time.Now()
	traceId := req.GetTraceId()

	defer func() {
		elapsed := time.Since(startTime).Microseconds()
		monitoring.RecordMetrics("MarkImpression", elapsed)
		log.Info().Str("Component", "MarkImpression").Str("TraceId:", traceId).Msgf("MarkImpression req %v took %v microseconds", req.GetSkuId(), elapsed)
	}()

	skuId := req.GetSkuId()
	userId := req.GetUserId()
	expiresAt, err := time.Parse(time.RFC3339, req.GetExpiresAt())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse expiresAt: %v", err)
	}

	err = s.m.MarkImpression(skuId, userId, expiresAt)

	return &pb.MarkImpressionResp{}, err
}

func serverInterceptor(m *inmemory.InMemoryStore) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		// Calls the handler
		h, err := handler(ctx, req)

		method := strings.Replace(info.FullMethod, "/"+pb.FeatureStore_ServiceDesc.ServiceName+"/", "", 1)
		if _, ok := constant.EXCLUDE_METHOD_BLOCK_GRPC[method]; !ok {
			if c := m.WaitReadyQ(); c != nil {
				log.Debug().Msg("blocking...")
				<-c
				log.Debug().Msg("block released!")
			}
		}

		monitoring.RecordMetrics(method, time.Since(start).Microseconds())

		return h, err
	}
}

func StartListener(mainCtx context.Context, m *inmemory.InMemoryStore, wg *sync.WaitGroup) {
	defer wg.Done()
	port := util.GetConfig().GrpcPort
	log.Info().Msgf("GRPC server starting at port: %v ", port)

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(serverInterceptor(m)))

	reflection.Register(s)

	grpcServer := server{
		m:      m,
		locker: util.NewLockManager(),
	}

	pb.RegisterFeatureStoreServer(s, &grpcServer)

	// Start serving in a goroutine, so it won't block the graceful shutdown handling below
	go func() {
		time.AfterFunc(1*time.Second, func() {
			api.STATUS.Add(1)
		})
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	<-mainCtx.Done()
	log.Info().Str("Component", "inmemory.StartListener").Msgf("Got signal. GRPC Server will shut down in %v seconds", api.WAIT_SECONDS)
	time.Sleep(time.Duration(api.WAIT_SECONDS) * time.Second)

	log.Info().Str("Component", "inmemory.StartListener").Msg("Gracefully stopping server...")

	// GracefulStop stops the gRPC server gracefully. It stops the server from
	// accepting new connections and RPCs and blocks until all the pending RPCs are finished.
	s.GracefulStop()

	log.Info().Str("Component", "inmemory.StartListener").Msg("Server has stopped")
}
