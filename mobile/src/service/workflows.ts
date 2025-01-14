export async function sendWorkflows(
  token: string,
  apiEndpoint: string,
  formsRegister: { action_id: number; reaction_id: number; name?: string },
) {
  try {
    const response = await fetch(`http://${apiEndpoint}:8080/api/workflow`, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      method: 'POST',
      body: JSON.stringify({
        action_id: formsRegister.action_id,
        reaction_id: formsRegister.reaction_id,
        name: formsRegister.name,
      }),
    });
    if (response.status !== 200) console.error('API send Workflows error');
    return true;
  } catch (error) {
    console.error('Error fetching workflows data:', error);
    return false;
  }
}

export async function getReaction(
  apiEndpoint: string,
  token: string,
  sendReaction: (reaction: any) => void,
) {
  console.log('apiEndpoint', apiEndpoint, token);
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/workflow/reaction`,
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
      console.log('data', data);
      if (data !== null) sendReaction(data);
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
          Authorization: `Bearer ${token}`,
        },
        method: 'GET',
      },
    );
    const data = await response.json();
    if (response.status == 200) {
      if (data !== null) setWorkflows(data);
    } else {
      console.error('Error invalide token');
    }
    return true;
  } catch (error) {
    console.error('Error fetching Workflows data:', error);
    return false;
  }
}

export async function modifyWorkflows(
  apiEndpoint: string,
  token: string,
  workflowStatus: boolean,
  workflowId: number,
) {
  console.log('workflowStatus', workflowStatus);
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/workflow/activation`,
      {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        method: 'PUT',
        body: JSON.stringify({
          workflow_id: workflowId,
          workflow_state: workflowStatus,
        }),
      },
    );
    if (response.status !== 200) console.error('Error invalide token');
    return true;
  } catch (error) {
    console.error('Error put Workflows data:', error);
    return false;
  }
}

export async function deleteWorkflow(
  apiEndpoint: string,
  token: string,
  workflowId: number,
  workflowName: string,
  actionId: number,
  reactionId: number,
) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/workflow`,
      {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        method: 'DELETE',
        body: JSON.stringify({
          workflow_id: workflowId,
          name: workflowName,
          action_id: actionId,
          reaction_id: reactionId,
        }),
      },
    );
    if (response.status !== 200) console.error('Error invalide token');
    return true;
  } catch (error) {
    console.error('Error put Workflows data:', error);
    return false;
  }
}