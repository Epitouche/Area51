export async function sendWorkflows(
  token: string,
  apiEndpoint: string,
  formsRegister: { action_id: number; reaction_id: number },
) {
  try {
    const response = await fetch(`http://${apiEndpoint}:8080/api/workflow`, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      method: 'POST',
      body: JSON.stringify({
        action_id: formsRegister.action_id,
        reaction_id: formsRegister.reaction_id,
      }),
    });
    if (response.status === 200) {
      console.log('API send Workflows success');
    } else {
      console.error('API send Workflows error:', response.status, response.statusText, token);
  }
    return true;
  } catch (error) {
    console.error('Error fetching workflows data:', error);
    return false;
  }
}

export async function getWorkflows(apiEndpoint: string, token: string, sendWorkflows: (workflow: any) => void) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/workflow/reaction`,
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
      console.log('API get workflows success');
      if (data !== null)
        sendWorkflows(data);
    }
    return true;
  } catch (error) {
    console.error('Error fetching Workflows data:', error);
    return false;
  }
}