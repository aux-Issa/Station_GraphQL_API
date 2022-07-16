import {GraphQLClient} from "graphql-request";
import {getSdk} from "./src/graphql/generated/graphql";

async function main() {
  const client = new GraphQLClient("http://localhost:8080/graphql")
  const sdk = getSdk(client)
  const res = await sdk.stationByCD({
      id: 1130201,
  })
  console.log(res)
}

main()