import { Construct } from 'constructs';
import { App, TerraformStack } from 'cdktf';
import { Team, ConcourseProvider } from './.gen/providers/concourse'

class MyStack extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);

    new ConcourseProvider(this, 'concourse', {
      url: process.env.CONCOURSE_URL as string,
      username: process.env.CONCOURSE_USERNAME as string,
      password: process.env.CONCOURSE_PASSWORD as string,
    })

    new Team(this, 'testTeam', {
      name: "test-team",
      ownerUsers: ["local:cappyzawa"],
    })
  }
}

const app = new App();
new MyStack(app, 'cdk');
app.synth();
