# Python APIの自動生成方法

OpenAPI Generator v6.6.0を用いる。現状の最新版(v7.0.0)ではバイナリをうまく取り扱えない。

https://github.com/OpenAPITools/openapi-generator

## 生成

まず、setup.pyとLICENSEをどこかに移動させる。

integration.ymlのあるディレクトリで

```
$ docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate \
  -i /local/integration.yml \
  -g python \
  -o /local/python --package-name reearthcmsapi
```

を実行すると、pythonディレクトリにコードが生成される。

## 修正

`reearthcmsapi/configuration.py`にバグがあるため、以下の箇所を

```
self.access_token = None
```

以下のように変更する。

```
self.access_token = access_token
```

コードをコピーし、さらに先ほど移動させたsetup.pyとLICENSEをコピーする。setup.pyのバージョンの内容などは適宜変更する。

## リリース

まず形式のチェックを行う。

```
pip install wheel twine
python setup.py sdist
python setup.py bdist_wheel
twine check dist/*
```

以下を参考にリリースを行う。

https://packaging.python.org/en/latest/tutorials/packaging-projects/