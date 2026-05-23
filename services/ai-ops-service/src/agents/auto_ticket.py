import json
import requests

class AutoTicketAgent:
    def __init__(self, incident_svc_url: str):
        self.incident_svc_url = incident_svc_url

    def create_ticket_from_alarm(self, alarm_event: dict):
        """
        AI-driven logic to determine if an alarm requires an automated trouble ticket.
        """
        severity = alarm_event.get('severity')
        resource_id = alarm_event.get('resource_id')

        if severity in ['Critical', 'Major']:
            print(f"AI: Alarm on {resource_id} is {severity}. Creating automated trouble ticket.")

            ticket_data = {
                "description": f"Automated ticket for {severity} alarm on {resource_id}",
                "severity": severity,
                "status": "Open",
                "relatedParty": [{"id": "AI-OPS-01", "role": "Originator"}]
            }

            # POST to incident-service
            # requests.post(f"{self.incident_svc_url}/troubleTicket", json=ticket_data)
            return True
        return False
