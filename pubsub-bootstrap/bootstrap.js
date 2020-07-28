const { PubSub } = require("@google-cloud/pubsub");
const pubsub = new PubSub();

const init = async () => {
  const [topic] = await pubsub
    .topic("requests").create();
  await topic
    .createSubscription("api", {
      pushConfig: {
        pushEndpoint: "http://api/event",
      },
    });  
};

init().catch(console.log);
