import { GraphQLClient } from 'graphql-request';
import * as Dom from 'graphql-request/dist/types.dom';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Query = {
  __typename?: 'Query';
  stationByCD: Station;
  stationByName?: Maybe<Array<Maybe<Station>>>;
};


export type QueryStationByCdArgs = {
  stationCD?: InputMaybe<Scalars['Int']>;
};


export type QueryStationByNameArgs = {
  stationName?: InputMaybe<Scalars['String']>;
};

export type Station = {
  __typename?: 'Station';
  address?: Maybe<Scalars['String']>;
  afterStation?: Maybe<Station>;
  beforeStation?: Maybe<Station>;
  lineName?: Maybe<Scalars['String']>;
  stationCD: Scalars['Int'];
  stationName: Scalars['String'];
  transferStation?: Maybe<Array<Maybe<Station>>>;
};

export type StationByCdQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type StationByCdQuery = { __typename?: 'Query', stationByCD: { __typename?: 'Station', lineName?: string | null, stationCD: number, stationName: string, beforeStation?: { __typename?: 'Station', lineName?: string | null, stationCD: number, stationName: string, transferStation?: Array<{ __typename?: 'Station', lineName?: string | null, stationCD: number, stationName: string } | null> | null } | null, afterStation?: { __typename?: 'Station', lineName?: string | null, stationCD: number, stationName: string } | null } };


export const StationByCdDocument = gql`
    query stationByCD($id: Int!) {
  stationByCD(stationCD: $id) {
    lineName
    stationCD
    stationName
    beforeStation {
      lineName
      stationCD
      stationName
      transferStation {
        lineName
        stationCD
        stationName
      }
    }
    afterStation {
      lineName
      stationCD
      stationName
    }
  }
}
    `;

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string, operationType?: string) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName, _operationType) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    stationByCD(variables: StationByCdQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<StationByCdQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<StationByCdQuery>(StationByCdDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'stationByCD', 'query');
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;