import { AboutJson } from '../types';

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
    if (response.status === 200) {
      setAboutJson(data);
    }
    return data;
  } catch (error) {
    console.error('Error fetching AboutJson data:', error, apiEndpoint);
  }
}
