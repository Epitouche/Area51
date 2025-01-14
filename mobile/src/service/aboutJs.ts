import {
  AboutJson,
  AboutJsonParse,
  ConnectedService,
  GetConnectedServiceProps,
  ParseConnectedServicesProps,
  ParseServicesProps,
} from '../types';
import { getToken } from './token';

export async function getAboutJson(
  apiEndpoint: string,
  setAboutJson: (aboutjson: AboutJson) => void,
) {
  try {
    const response = await fetch(`http://${apiEndpoint}:8080/about.json`, {
      method: 'GET',
      body: null,
    });
    const data = await response.json();
    if (response.status === 200) setAboutJson(data);
    return data;
  } catch (error) {
    console.error('Error fetching AboutJson data:', error);
  }
}

export async function parseServices({
  aboutJson,
  serverIp,
  setServicesConnected,
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
  let aboutJsonParse: AboutJsonParse = {
    services: [],
  };
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

  aboutJsonParse = {
    services: aboutjson.server.services.map(service => {
      const connected = connectedServices.some(
        connectedService => connectedService.name === service.name,
      );
      return {
        name: service.name,
        isConnected: connected,
        actions: service.actions,
        reactions: service.reactions,
        image: service.image,
      };
    }),
  };
  if (aboutJsonParse.services.length > 0) setServicesConnected(aboutJsonParse);
}

export async function getConnectedService({
  apiEndpoint,
  token,
  setConnectedService,
}: GetConnectedServiceProps) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/user/services `,
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
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
