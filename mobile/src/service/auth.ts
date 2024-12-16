import { saveToken } from './token';
import { LoginProps, RegisterProps } from '../types';
import { useEffect, useState } from 'react';

interface AuthApiCall {
  apiEndpoint: string;
  formsLogin?: LoginProps;
  formsRegister?: RegisterProps;
  setMessage: (message: string) => void;
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
        setMessage('Error: invalid credentials');
        return false;
      }
      setMessage('Error: token not found');
      console.error('Token not found');
      return false;
    }
    await saveToken('token', data.access_token);
    console.log('API login success');
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
    await saveToken('token', data.access_token);
    return true;
  } catch (error) {
    setMessage('Error: Internal Server Error');
    console.error('API call failed:', error);
    return false;
  }
}
