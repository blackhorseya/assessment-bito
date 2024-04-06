import http from 'k6/http';
import {check, group} from 'k6';
import {uuidv4} from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export const options = {
  vus: 1,
  iterations: 1,
};

const BASE_URL = 'http://localhost:1992';

export default function() {
  let url = `${BASE_URL}/api/v1/players`;

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const tests = [
    {
      name: 'RemoveSinglePerson - invalid id then 400',
      id: 'invalid',
      expectedStatus: 400,
    },
    {
      name: 'RemoveSinglePerson - not found id then 404',
      id: uuidv4(),
      expectedStatus: 500,
    },
  ];

  for (let test of tests) {
    group(test.name, () => {
      const resp = http.del(`${url}/${test.id}`, null, params);

      if (check(resp, {
        [`status equals ${test.expectedStatus}`]: (resp) => resp.status ===
            test.expectedStatus,
      })) {
        console.debug(`Added player and match: ${resp.status} ${resp.body}`);
      } else {
        console.error(
            `Failed to add player and match: ${resp.status} ${resp.body}`);
      }
    });
  }
}
