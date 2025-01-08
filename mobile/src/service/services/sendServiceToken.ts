import { deleteToken, saveToken } from "../token";

export async function sendServiceToken(
  apiEndpoint: string,
  token: string,
  name: string,
) {
  try {
    let response;
      response = await fetch(
        `http://${apiEndpoint}:8080/api/github/mobile/token`,
        {
          method: 'POST',
          body: JSON.stringify({ token: token, service: name }),
        },
      );
    const data = await response.json();
    if (response.status === 200) {
      console.log('API send Github Token success');
    }
    deleteToken('token');
    saveToken('token', data.token);
    return true;
  } catch (error) {
    console.error('Error:', error);
    return false;
  }
}
