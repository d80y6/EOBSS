import json
from kafka import KafkaConsumer
from .rca_engine import RCAEngine

class AIKafkaConsumer:
    def __init__(self, brokers: list, topic: str):
        self.consumer = KafkaConsumer(
            topic,
            bootstrap_servers=brokers,
            value_deserializer=lambda m: json.loads(m.decode('utf-8'))
        )
        self.rca_engine = RCAEngine()

    def start_listening(self):
        print(f"AI Ops listening on Kafka topic...")
        for message in self.consumer:
            event = message.value
            # Perform real-time anomaly detection
            if event.get('type') == 'ALARM':
                self.rca_engine.analyze_alarm(event)
