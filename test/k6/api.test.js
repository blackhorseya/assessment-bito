import {group} from 'k6';

export const options = {
  vus: 1,
  iterations: 1,

  // The following section contains configuration options for execution of this
  // test script in Grafana Cloud.
  //
  // See https://grafana.com/docs/grafana-cloud/k6/get-started/run-cloud-tests-from-the-cli/
  // to learn about authoring and running k6 test scripts in Grafana k6 Cloud.
  //
  // ext: {
  //   loadimpact: {
  //     // The ID of the project to which the test is assigned in the k6 Cloud UI.
  //     // By default tests are executed in default project.
  //     projectID: "",
  //     // The name of the test in the k6 Cloud UI.
  //     // Test runs with the same name will be grouped.
  //     name: "api.test.js"
  //   }
  // },
};

const BASE_URL = 'http://localhost:1992/api';

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function() {
  group('/v1/players', () => {
    // todo: 2024/4/8|sean|implement this test
  });

  group('/v1/players/{id}', () => {
    // todo: 2024/4/8|sean|implement this test
  });

  group('/v1/pairs', () => {
    // todo: 2024/4/8|sean|implement this test
  });
}
