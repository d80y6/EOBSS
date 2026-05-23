import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 50 }, // ramp up to 50 users
    { duration: '1m', target: 50 },  // stay at 50 users
    { duration: '20s', target: 0 },  // scale down
  ],
};

export default function () {
  const url = 'http://order-service.telcoflow.svc.cluster.local:8080/productOrder';
  const payload = JSON.stringify({
    externalId: \`EXT-\${__ITER}\`,
    customer: { id: "CUST-001" },
    orderItem: [
      { action: "add", productOffering: { id: "OFF-001" } }
    ]
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'X-Tenant-ID': 'TENANT-01',
    },
  };

  const res = http.post(url, payload, params);
  check(res, {
    'status is 201': (r) => r.status === 201,
    'transaction time < 200ms': (r) => r.timings.duration < 200,
  });
  sleep(1);
}
