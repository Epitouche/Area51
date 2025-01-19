// src/services/api.js

import { deleteToken, saveToken } from "./token";

export async function sendGithubToken(
  apiEndpoint: string,
  token: string,
) {
  try {
    const response = await fetch(`http://${apiEndpoint}:8080/api/github/mobile/token`, {
      method: 'POST',
      body: JSON.stringify({'token': token}),
    });
    const data = await response.json();
    if (response.status === 200) {
      console.log('API send Github Token success:');
    }
    deleteToken('token');
    saveToken('token', data.token);
    return true;
  } catch (error) {
    console.error('Error:', error, apiEndpoint);
    return false;
  }
}
