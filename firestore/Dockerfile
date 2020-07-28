FROM google/cloud-sdk

RUN apt-get install -y google-cloud-sdk-firestore-emulator

ENTRYPOINT [ "gcloud", "alpha", "emulators", "firestore", "start" ]


