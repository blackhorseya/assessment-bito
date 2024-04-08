import {check, group} from 'k6';
import http from 'k6/http';
import {
  randomIntBetween,
  randomString,
} from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export const options = {
  // The following section contains configuration options for execution of this
  // test script in Grafana Cloud.
  //
  // See https://grafana.com/docs/grafana-cloud/k6/get-started/run-cloud-tests-from-the-cli/
  // to learn about authoring and running k6 test scripts in Grafana k6 Cloud.
  //
  ext: {
    loadimpact: {
      // The ID of the project to which the test is assigned in the k6 Cloud UI.
      // By default tests are executed in default project.
      projectID: '3690299',
      // The name of the test in the k6 Cloud UI.
      // Test runs with the same name will be grouped.
      name: 'api.e2e.js',
    },
  },
};

const BASE_URL = 'http://localhost:1992/api';

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function() {
  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  group('/v1/players', () => {
    let url = `${BASE_URL}/v1/players`;

    let player1 = {
      id: '',
      age: randomIntBetween(1, 100),
      gender: randomIntBetween(1, 2),
      height: randomIntBetween(100, 200),
      name: randomString(10),
      nums_of_wanted_dates: randomIntBetween(1, 10),
    };

    {
      const resp = http.post(url, JSON.stringify(player1), params);

      if (check(resp,
          {'create a player': (r) => r.status === 201})) {
        player1.id = resp.json().data.id;
        console.debug(`Player created successfully, ID: ${player1.id}`);
      } else {
        console.error(
            `Failed to add player and match: ${resp.status} ${resp.body}`);
      }
    }

    {
      const resp = http.get(url);

      if (check(resp,
          {'list players': (r) => r.status === 200})) {
        console.debug(
            `Players retrieved successfully: ${resp.headers['X-Total-Count']}`);
      } else {
        console.error(
            `Failed to retrieve players: ${resp.status} ${resp.body}`);
      }
    }

    {
      const resp = http.get(`${url}/${player1.id}`);

      if (check(resp,
          {'get player by id': (r) => r.status === 200})) {
        console.debug(`Player retrieved successfully: ${resp.json().data.id}`);
      } else {
        console.error(
            `Failed to retrieve player: ${resp.status} ${resp.body}`);
      }
    }

    {
      const resp = http.del(`${url}/${player1.id}`);

      if (check(resp,
          {'delete player by id': (r) => r.status === 204})) {
        console.debug(`Player deleted successfully: ${player1.id}`);
      } else {
        console.error(
            `Failed to delete player: ${resp.status} ${resp.body}`);
      }
    }
  });

  group('/v1/pairs', () => {
    let left = {
      id: '',
      age: randomIntBetween(1, 100),
      gender: 1,
      height: randomIntBetween(170, 200),
      name: randomString(10),
      nums_of_wanted_dates: randomIntBetween(1, 10),
    };

    let right = {
      id: '',
      age: randomIntBetween(1, 100),
      gender: 2,
      height: randomIntBetween(140, 170),
      name: randomString(10),
      nums_of_wanted_dates: randomIntBetween(1, 10),
    };
    let players = [left, right];

    {
      players.forEach((player) => {
        let url = `${BASE_URL}/v1/players`;
        const resp = http.post(url, JSON.stringify(player), params);

        if (check(resp,
            {'create a player': (r) => r.status === 201})) {
          player.id = resp.json().data.id;
          console.debug(`Player created successfully, ID: ${player.id}`);
        } else {
          console.error(
              `Failed to add player and match: ${resp.status} ${resp.body}`);
        }
      });
    }

    {
      let url = `${BASE_URL}/v1/pairs`;
      let body = {
        left_id: left.id,
        right_id: right.id,
      };
      const resp = http.post(url, JSON.stringify(body), params);

      if (check(resp,
          {'create a pair': (r) => r.status === 201})) {
        console.debug(`Pair created successfully: ${resp.body}`);
      } else {
        console.error(
            `Failed to add pair: ${resp.status} ${resp.body}`);
      }
    }
  });
}
