import { useEffect, useState } from 'react';
import { AboutJson, AboutJsonParse, ConnectedService } from '../types';
import { getToken } from './token';

type Workflows = {
  apiEndpoint: string;
  token: string;
  setConnectedService: (connectedService: ConnectedService[]) => void;
};

interface ParseConnectedServicesProps {
  aboutjson: AboutJson;
  apiEndpoint: string;
  token: string;
  setServicesConnected: (servicesConnected: AboutJsonParse) => void;
}

interface ParseServicesProps {
  aboutJson: AboutJson;
  serverIp: string;
  setServicesConnected: (servicesConnected: AboutJsonParse) => void;
}

export async function parseServices({
  aboutJson,
  serverIp,
  setServicesConnected
}: ParseServicesProps) {
  let myToken: string = '';
  const setToken = (token: string) => {
    myToken = token;
  };
  await getToken('token', setToken);
  if (myToken !== 'Error: token not found' && myToken !== '' && aboutJson) {
    parseConnectedServices({
      token: myToken,
      aboutjson: aboutJson,
      apiEndpoint: serverIp,
      setServicesConnected,
    });
  }
}

export async function parseConnectedServices({
  aboutjson,
  apiEndpoint,
  token,
  setServicesConnected,
}: ParseConnectedServicesProps) {
  let connectedServices: ConnectedService[] = [];
  const setConnectedService = (connectedService: ConnectedService[]) => {
    connectedServices = connectedService;
  };
  const connectedServicesCall = async () => {
    if (token !== 'Error: token not found') {
      await getConnectedService({
        apiEndpoint,
        token,
        setConnectedService,
      });
    }
  };
  await connectedServicesCall();

  const aboutJsonParse: AboutJsonParse = {
    services: aboutjson.server.services.map(service => {
      const connected = connectedServices.some(
        connectedService => connectedService.name === service.name,
      );
      return {
        name: service.name,
        isConnected: connected,
        actions: service.actions,
        reactions: service.reactions,
      };
    }),
  };
  setServicesConnected(aboutJsonParse);
}



export async function getConnectedService({
  apiEndpoint,
  token,
  setConnectedService,
}: Workflows) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/user/services `,
      {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        method: 'GET',
      },
    );
    const data = await response.json();
    if (response.status === 200) {
      setConnectedService(data);
    } else if (response.status !== 200) {
      console.error('Error fetching services data');
      setConnectedService([
        {
          created_at: '',
          description: '',
          id: 0,
          name: '',
          updated_at: '',
        },
      ]);
    }
    return true;
  } catch (error) {
    console.error('Error fetching services data:', error);
    return false;
  }
}
