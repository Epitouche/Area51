export async function sendWorkflows(
  token: string,
  apiEndpoint: string,
  formsRegister: { action_id: number; reaction_id: number; name?: string },
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
        name: formsRegister.name,
      }),
    });
    if (response.status === 200) {
      console.log('API send Workflows success');
    } else {
      console.error('API send Workflows error',);
  }
    return true;
  } catch (error) {
    console.error('Error fetching workflows data:', error);
    return false;
  }
}

export async function getReaction(apiEndpoint: string, token: string, sendReaction: (reaction: any) => void) {
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
      if (data !== null)
        sendReaction(data);
    }
    return true;
  } catch (error) {
    console.error('Error fetching Reaction data:', error);
    return false;
  }
}

export async function getWorkflows(
  apiEndpoint: string,
  token: string,
  setWorkflows: (workflows: any) => void,
) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/user/workflows`,
      {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        method: 'GET',
      },
    );
    const data = await response.json();
    if (response.status === 200) {
      if (data !== null)
        setWorkflows(data);
    } else {
      console.error('Error invalide token');
    }
    return true;
  } catch (error) {
    console.error('Error fetching Workflows data:', error);
    return false;
  }
}

