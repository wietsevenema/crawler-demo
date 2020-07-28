<?php

use Google\Cloud\Firestore\FirestoreClient;

require 'vendor/autoload.php';

$firestoreClient = new FirestoreClient(
    [
        'projectId' => getenv('GOOGLE_PROJECT'),
    ]
);

foreach ($firestoreClient->collection('requests')->listDocuments() as $document) {
    print_r($document->snapshot()->data());
}
