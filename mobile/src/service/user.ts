import { deleteToken, saveToken } from './token';
import { LoginProps, RegisterProps } from '../types';

interface AuthApiCall {
  apiEndpoint: string;
  formsLogin?: LoginProps;
  formsRegister?: RegisterProps;
  setMessage: (message: string) => void;
}

interface DeleteUser {
  apiEndpoint: string;
  token: string;
}

export async function loginApiCall({
  apiEndpoint,
  formsLogin,
  setMessage,
}: AuthApiCall) {
  try {
    if (!formsLogin) {
      throw new Error('No login form provided');
    }
    const formData = new FormData();
    formData.append('username', formsLogin.username);
    formData.append('password', formsLogin.password);
    const response = await fetch(`http://${apiEndpoint}:8080/api/auth/login`, {
      method: 'POST',
      body: formData,
    });
    const data = await response.json();
    if (response.status !== 200) {
      if (response.status === 401) {
        setMessage('Username or password is incorrect');
        return false;
      }
      setMessage('Error: token not found');
      console.error('Token not found');
      return false;
    }
    await saveToken('token', data.access_token);
    return true;
  } catch (error) {
    setMessage('Error: Internal Server Error');
    console.error('API call failed:', error);
    return false;
  }
}

export async function registerApiCall({
  apiEndpoint,
  setMessage,
  formsRegister,
}: AuthApiCall) {
  try {
    if (!formsRegister) {
      throw new Error('No register form provided');
    }
    const formData = new FormData();
    formData.append('email', formsRegister.email);
    formData.append('password', formsRegister.password);
    formData.append('username', formsRegister.username);
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/auth/register`,
      {
        method: 'POST',
        body: formData,
      },
    );
    const data = await response.json();
    if (response.status !== 200) {
      if (response.status === 409) {
        setMessage('Error: email already exists');
        return false;
      }
      setMessage('Error: token not found');
      console.error('Token not found');
      return false;
    }
    await deleteToken('token');
    await saveToken('token', data.access_token);
    return true;
  } catch (error) {
    setMessage('Error: Internal Server Error');
    console.error('API call failed:', error);
    return false;
  }
}

export async function deleteUser({ apiEndpoint, token }: DeleteUser) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/user/account`,
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        method: 'DELETE',
      },
    );
    if (response.status !== 200) {
      console.error('Error fetching user data');
      return false;
    }
    return true;
  } catch (error) {
    console.error('API call failed:', error);
    return false;
  }
}
