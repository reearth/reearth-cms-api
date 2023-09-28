# How to automatically generate Python API

We use OpenAPI Generator v6.6.0 to generate the Python API using integration.yml. We currently do not use v7.0.0 because it cannot handle binary files described in integration.yml.  If integration.yml is changed, we need to generate new API source code.

https://github.com/OpenAPITools/openapi-generator

## Preparation

First, we have modified and added some files from the original generated source code.  So we need to make a backup of these files:

```
setup.py
LICENSE
requirements.txt
```

## Code Generation

Run this command in the directory that contains integration.yml.

```
$ docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate \
  -i /local/integration.yml \
  -g python \
  -o /local/python --package-name reearthcmsapi
```

The source code will be generated in the `python` directory.  Copy it to this repository.

## Modify bug

Since `reearthcmsapi/configuration.py` has a bug, we need to modify this line

```
self.access_token = None
```

to this.

```
self.access_token = access_token
```

## Copying files

Then, copy the backup files to the `python` directory.

## Release

To release, we need to check that the project structure is valid.

```
pip install wheel twine
python setup.py sdist
python setup.py bdist_wheel
twine check dist/*
```

Then publish the package by following this URL.

https://packaging.python.org/en/latest/tutorials/packaging-projects/