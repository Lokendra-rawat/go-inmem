package inmemory

import (
	"time"

	pb "example.com/go-inmem-loki/proto"
	"example.com/go-inmem-loki/types"
	"example.com/go-inmem-loki/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type InMemoryStore struct {
	env         string
	skuFeatures *util.TTLCache[types.InMemorySkuFeatures]
	existsMap   *util.TTLCache[bool]
	readyQ      *util.ReadyQueue
	pool        *pgxpool.Pool
	podId       string
	bootupTime  time.Time
}

// func Init(ctx context.Context, wg *sync.WaitGroup) *InMemoryStore {
// 	defer wg.Done()
// 	m := &InMemoryStore{}

// 	m.pool = db.GetDB()
// 	m.bootupTime = time.Now()

// 	m.skuFeatures = util.NewTTLCache[types.InMemorySkuFeatures](24 * time.Hour)
// 	m.existsMap = util.NewTTLCache[bool](24 * time.Hour)

// 	m.env = util.GetConfig().GoEnv

// 	if m.env != "test" {
// 		wg.Add(1)
// 		go kafka.DebugDelivery(ctx, wg)
// 		wg.Add(1)
// 		go kafka.FlushMessagesPeriodically(ctx, wg)
// 	}

// 	m.podId = os.Getenv("HOSTNAME")
// 	if m.podId == "" {
// 		m.podId = uuid.NewString()
// 	}
// 	m.readyQ = util.NewReadyQueue(m.podId)

// 	return m
// }

func (m *InMemoryStore) GetAllSkuFeatures() map[string]util.TTLCacheData[types.InMemorySkuFeatures] {
	return m.skuFeatures.Dump()
}

func (m *InMemoryStore) GetSkuFeatures(key string, catalogueName string, traceId string) *pb.SkuFeatures {
	f, ok := m.skuFeatures.Get(key)
	if !ok {
		log.Warn().Str("Component", "InMemoryStore.GetSkuFeatures").Str("TraceId", traceId).Str("key", key).Msg("key not found in cache")
		return nil
	}

	return util.ConvertInmemoryToPbSkuFeatures(f.Value, catalogueName)
}

func (m *InMemoryStore) SetSkuFeatures(key string, value types.InMemorySkuFeatures, expiry *time.Duration) {
	m.skuFeatures.Set(key, value, expiry)
}

// func (m *InMemoryStore) GetSkuFeaturesForAllCatalogues(skuId string, traceId string) map[string]*pb.SkuFeatures {
// 	f, ok := m.skuFeatures.Get(skuId)
// 	if !ok {
// 		log.Warn().Str("Component", "InMemoryStore.GetSkuFeaturesForAllCatalogues").Str("TraceId", traceId).Str("key", skuId).Msg("key not found in cache")
// 		return nil
// 	}

// 	return util.ConvertInmemoryToPbSkuFeaturesForAllCatalogues(f.Value)
// }

// func (m *InMemoryStore) HandleSkuUpdatePayload(payload types.UpdateSkuPayload) {

// 	val := types.InMemorySkuFeatures{}
// 	f, ok := m.skuFeatures.Get(payload.SkuId)
// 	if ok {
// 		val = f.Value
// 	}

// 	// Optional Expiry
// 	var expiry *time.Duration
// 	if payload.TTLSeconds != nil {
// 		ttl := time.Duration(*payload.TTLSeconds) * time.Second
// 		expiry = &ttl
// 	}

// 	util.CovertPbToInmemorySkuFeatures(payload.Features, &val, payload.CatalogueName)
// 	m.skuFeatures.Set(payload.SkuId, val, expiry)

// }

// func (m *InMemoryStore) MarkAsExisting(key string, expiry time.Duration) {
// 	m.existsMap.Set(key, true, &expiry)
// }

// func (m *InMemoryStore) UpsertSkuFeatures(payload types.UpdateSkuPayload) error {
// 	m.HandleSkuUpdatePayload(payload)

// 	var err error
// 	if m.env != "test" {
// 		err = kafka.SendMessageToKafka(kafka.RankerFeatureStoreUpdatesTopic, types.KafkaCommand[types.UpdateSkuPayload]{
// 			Type:    constant.EventSkuFeaturesUpsert,
// 			Payload: payload,
// 			PodId:   m.podId,
// 		})
// 		if err != nil {
// 			log.Err(err).Msgf("Error sending message to kafka struct %+v", payload)
// 			panic(err)
// 		}
// 	}

// 	return err
// }

// func (m *InMemoryStore) OrderExists(skuId string, userId int64) bool {
// 	key := OrderBoolKey(skuId, userId)

// 	f, ok := m.existsMap.Get(key)
// 	if !ok {
// 		return false
// 	}

// 	return f.Value
// }

// func (m *InMemoryStore) MarkOrder(skuId string, userId int64, expiresAt time.Time) error {
// 	key := OrderBoolKey(skuId, userId)

// 	duration := time.Until(expiresAt)
// 	m.existsMap.Set(key, true, &duration)

// 	var err error
// 	if m.env != "test" {
// 		err = kafka.SendMessageToKafka(kafka.RankerFeatureStoreUpdatesTopic, types.KafkaCommand[types.UpdateCountPayload]{
// 			Type: constant.EventExistUpdate,
// 			Payload: types.UpdateCountPayload{
// 				Key:       key,
// 				ExpiresAt: expiresAt.Format(time.RFC3339),
// 			},
// 			PodId: m.podId,
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	return err
// }

// func (m *InMemoryStore) GetImpressionExists(skuId string, userId int64) bool {
// 	key := getImpressionBoolKey(skuId, userId)

// 	f, ok := m.existsMap.Get(key)
// 	if !ok {
// 		return false
// 	}

// 	return f.Value
// }

// func (m *InMemoryStore) MarkImpression(skuId string, userId int64, expiresAt time.Time) error {
// 	key := getImpressionBoolKey(skuId, userId)

// 	duration := time.Until(expiresAt)
// 	m.existsMap.Set(key, true, &duration)

// 	var err error
// 	if m.env != "test" {
// 		err = kafka.SendMessageToKafka(kafka.RankerFeatureStoreUpdatesTopic, types.KafkaCommand[types.UpdateCountPayload]{
// 			Type: constant.EventExistUpdate,
// 			Payload: types.UpdateCountPayload{
// 				Key:       key,
// 				ExpiresAt: expiresAt.Format(time.RFC3339),
// 			},
// 			PodId: m.podId,
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	return err
// }

// func (m *InMemoryStore) WaitReadyQ() chan struct{} {
// 	return m.readyQ.Wait()
// }

// func (m *InMemoryStore) GetPodId() string {
// 	return m.podId
// }

// func (m *InMemoryStore) GetBootupTime() time.Time {
// 	return m.bootupTime
// }

// func (m *InMemoryStore) PopulationComplete() {
// 	m.readyQ.Done()
// }

// func getImpressionBoolKey(skuId string, userId int64) string {
// 	return "sku_id:" + skuId + ":user_id:" + cast.ToString(userId) + constant.IMPR_BOOL_SUFFIX
// }

// func OrderBoolKey(skuId string, userId int64) string {
// 	return "sku_id:" + skuId + ":user_id:" + cast.ToString(userId) + constant.ORDER_BOOL_SUFFIX
// }
