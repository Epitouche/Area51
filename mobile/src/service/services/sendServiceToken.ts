import { deleteToken, saveToken } from '../token';

export async function sendServiceToken(
  apiEndpoint: string,
  serviceToken: string,
  name: string,
  sessionToken?: string,
) {
  console.log(
    'serviceToken:',
    serviceToken,
    'name:',
    `"${name}"`,
    'sessionToken:',
    sessionToken,
  );
  try {
    let response;
    if (sessionToken === undefined) {
      console.log('API send Github Token without sessionToken');
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
    if (response.status === 200) {
      console.log('API send Github Token success');
    }
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
