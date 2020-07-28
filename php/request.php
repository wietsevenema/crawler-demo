<?php

use Google\Cloud\Firestore\FirestoreClient;
use Google\Cloud\PubSub\Message;
use Google\Cloud\PubSub\PubSubClient;

require 'vendor/autoload.php';

$project = getenv('GOOGLE_PROJECT');

$firestoreClient = new FirestoreClient(
    [
        'projectId' => $project,
    ]
);

$pubsubClient = new PubSubClient(
    [
        'projectId' => $project,
    ]
);

$request = ['url' => 'www.nos.nl'];

$document = $firestoreClient->collection('requests')->add($request);
$data = json_encode($document->snapshot()->data());
$pubsubClient->topic('requests')->publish(
    new Message(
        [
            'data' => $data,
        ]
    )
);
