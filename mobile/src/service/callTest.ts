// src/services/api.js

export async function exempleApiCall(
  apiEndpoint: string,
  setIsWorking: (isWorking: boolean) => void,
) {
  try {
    const response = await fetch(`http://${apiEndpoint}:8080/about.json`, {
      method: 'GET',
      body: null,
    });
    const data = await response.json();
    if (response.status === 200) {
    }
    setIsWorking(true);
    return data;
  } catch (error) {
    console.error('Error fetching user data:', error, apiEndpoint);
  }
}
