参考: https://user-first.ikyu.co.jp/entry/2019/12/16/113332

環境構築
```npx create-react-app frontend --template typescript```
```cd frontend```
```yarn add graphql-request```
```yarn add -D graphql @graphql-codegen/cli @graphql-codegen/typescript @graphql-codegen/typescript-graphql-request @graphql-codegen/typescript-operations```
frontend/graphql/queriesで定義したクエリファイルtsで生成する
```yarn run graphql-codegen --config codegen.yml```
CORS これ見る
https://qiita.com/Y_uuu/items/54c2749eca5b67ae9583