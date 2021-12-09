# controllermesh-api

Schema of the API types that are served by ControllerMesh.

## Purpose

This library is the canonical location of the ControllerMesh API definition.

We recommend using the go types in this repo. You may serialize them directly to JSON.

## Compatibility matrix

| ---------------------------------- |---------------------------|
| Kubernetes Version in your Project | Import controllermesh-api |
| ---------------------------------- |---------------------------| 
| < 1.18                             | v0.x.y-legacy             |
| >= 1.18                            | v0.x.y                    |

## Where does it come from?

`controllermesh-api` is synced from [https://github.com/openkruise/controllermesh/tree/master/apis](https://github.com/openkruise/controllermesh/tree/master/apis).
Code changes are made in that location, merged into `openkruise/kruise` and later synced here.

## Things you should NOT do

[https://github.com/openkruise/controllermesh/tree/master/apis](https://github.com/openkruise/controllermesh/tree/master/apis) is synced to here.
All changes must be made in the former. The latter is read-only.

