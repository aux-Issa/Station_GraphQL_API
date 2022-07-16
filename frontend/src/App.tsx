import React from 'react';
import logo from './logo.svg';
import './App.css';
import {GraphQLClient} from "graphql-request";
import {getSdk} from "./graphql/generated/graphql";

function App() {
  async function main() {
    const endpoint = "http://localhost:8080/graphql"
    const client = new GraphQLClient(endpoint,
      {headers: {
        authorization:'Access-Control-Allow-Origin'
      }}
      // headers: () => ({ 'X-Sent-At-Time': Date.now() })
    
    )
    const sdk = getSdk(client)
    const res = await sdk.stationByCD({
        id: 1130201,
    })
    debugger
    console.log(res)
  }
  
  main()
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
