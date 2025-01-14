import { deleteToken, saveToken } from '../token';
import { authorize } from 'react-native-app-auth';
import { githubLogin, microsoftLogin, spotifyLogin } from './index';
import { AuthApiCall, RefreshServicesProps, SelectServicesParamsProps } from '../../types';
import { getAboutJson, parseServices } from '../aboutJs';

export async function OauthLogin({ config, setToken }: AuthApiCall) {
  const authState = await authorize(config);
  if (!authState.accessToken) return false;
  setToken(authState.accessToken);
  return true;
}

export async function selectServicesParams({
  serverIp,
  serviceName,
  sessionToken,
}: SelectServicesParamsProps) {
  switch (serviceName) {
    case 'spotify':
      return await spotifyLogin(serverIp, sessionToken);
    case 'github':
      return await githubLogin(serverIp, sessionToken);
    case 'microsoft':
      return await microsoftLogin(serverIp, sessionToken);
    default:
      return false;
  }
}


export async function sendServiceToken(
  apiEndpoint: string,
  serviceToken: string,
  name: string,
  sessionToken?: string,
) {
  try {
    let response;
    if (sessionToken === undefined) {
      response = await fetch(`http://${apiEndpoint}:8080/api/mobile/token`, {
        headers: {
          'Content-Type': 'application/json',
        },
        method: 'POST',
        body: JSON.stringify({ token: serviceToken, service: name }),
      });
    } else {
      response = await fetch(`http://${apiEndpoint}:8080/api/mobile/token`, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${sessionToken}`,
        },
        method: 'POST',
        body: JSON.stringify({ token: serviceToken, service: name }),
      });
    }

    const data = await response.json();
    if (!sessionToken) {
      deleteToken('token');
      saveToken('token', data.token);
    }
    return true;
  } catch (error) {
    console.error('Error service OAuth2:', error);
    return false;
  }
}

export async function logoutServices(
  apiEndpoint: string,
  name: string,
  sessionToken: string,
) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/user/service/logout`,
      {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${sessionToken}`,
        },
        method: 'PUT',
        body: JSON.stringify({ service_name: name }),
      },
    );
    if (response.status === 200) {
      return true;
    }
    return false;
  } catch (error) {
    console.error('Error service OAuth2:', error);
    return false;
  }
}

export async function refreshServices({
  aboutJson,
  serverIp,
  setAboutJson,
  setServicesConnected,
}: RefreshServicesProps) {
  if (serverIp != '') {
    await getAboutJson(serverIp, setAboutJson);
    if (aboutJson)
      await parseServices({ serverIp, aboutJson, setServicesConnected });
  }
}

