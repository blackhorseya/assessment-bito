import http from 'k6/http';
import {check, group} from 'k6';

export const options = {
  scenarios: {
    my_scenario1: {
      executor: 'constant-arrival-rate',
      duration: '30s', // total duration
      preAllocatedVUs: 50, // to allocate runtime resources

      rate: 50, // number of constant iterations given `timeUnit`
      timeUnit: '1s',
    },
  },

  ext: {
    loadimpact: {
      projectID: 3690299,
      name: 'assessment-bito',
    },
  },
};

// Create a random string of given length
function randomString(length, charset = '') {
  if (!charset) charset = 'abcdefghijklmnopqrstuvwxyz';
  let res = '';
  while (length--) res += charset[(Math.random() * charset.length) | 0];
  return res;
}

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
      name: 'AddPlayerAndMatch then ok',
      payload: {
        age: 20,
        gender: 1,
        height: 180,
        name: randomString(10),
        nums_of_wanted_dates: 3,
      },
      expectedStatus: 201,
    },
    {
      name: 'AddPlayerAndMatch invalid age',
      payload: {
        age: -1,
        gender: 1,
        height: 180,
        name: randomString(10),
        nums_of_wanted_dates: 3,
      },
      expectedStatus: 400,
    },
    {
      name: 'AddPlayerAndMatch invalid height',
      payload: {
        age: 20,
        gender: 1,
        height: -180,
        name: randomString(10),
        nums_of_wanted_dates: 3,
      },
      expectedStatus: 400,
    },
    {
      name: 'AddPlayerAndMatch invalid nums_of_wanted_dates',
      payload: {
        age: 20,
        gender: 1,
        height: 180,
        name: randomString(10),
        nums_of_wanted_dates: -3,
      },
      expectedStatus: 400,
    },
  ];

  for (let test of tests) {
    group(test.name, () => {
      const resp = http.post(url, JSON.stringify(test.payload), params);

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
