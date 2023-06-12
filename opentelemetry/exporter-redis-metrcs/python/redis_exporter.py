from opentelemetry import metrics
from opentelemetry.metrics import CallbackOptions, Observation
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import PeriodicExportingMetricReader, ConsoleMetricExporter, MetricExporter
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.exporter.otlp.proto.http.metric_exporter import OTLPMetricExporter
from typing import Iterable
import redis
import os

class RedisAPI():
    connect = None
    def __init__(self, host, port) -> None:
        self.connect = redis.Redis(host=host, port=port)
    
    def getMemUsageByDB(self, db):
        self.connect.select(db)
        DBSize=0
        for k in self.connect.keys():
            DBSize += self.connect.memory_usage(k)
        print("DB {}: ".format(db),  DBSize)
    
    def getMemUsageByDBKeys(self, db):
        self.connect.select(db)
        key_mem_usage_list = {}
        for k in self.connect.keys():
            key_mem_usage_list[k] = self.connect.memory_usage(k)
        return list(sorted(key_mem_usage_list.items(), key=lambda x: x[1], reverse=True)[:10])
    
    def memory_usage_callback(self, options: CallbackOptions) -> Iterable[Observation]:
        temp = []
        for dbno in range(0, 15):
            temp.extend([Observation(i[1], {'key':i[0].decode("utf-8"), 'db': dbno}) for i in self.getMemUsageByDBKeys(dbno)])
        return temp

def main():
    REDIS_HOST=os.getenv('REDIS_HOST') or 'localhost'
    REDIS_PORT=os.getenv('REDIS_PORT') or 6379
    api = RedisAPI(REDIS_HOST, REDIS_PORT)
    resource = Resource(attributes={
        SERVICE_NAME: REDIS_HOST,
    })

    provider = MeterProvider(
        resource=resource,
        metric_readers=[PeriodicExportingMetricReader(OTLPMetricExporter(
        #endpoint="http://localhost:4318/v1/metrics"
            ))]
        )

    metrics.set_meter_provider(provider)

    meter = metrics.get_meter("redis")
    meter.create_observable_gauge(
        name="redis_keys_memory_usage_top10",
        callbacks=[api.memory_usage_callback],
        description="Returns redis key memory usage top 10",
    )


if __name__ == '__main__':
    main()
    