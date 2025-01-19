export async function sendWorkflows(
  token: string,
  apiEndpoint: string,
  formsRegister: {
    action_id: number;
    reaction_id: number;
    name?: string;
    action_option: any;
    reaction_option: any;
  },
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
        action_option: formsRegister.action_option,
        reaction_option: formsRegister.reaction_option,
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
  workflowId: number,
) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/workflow/reaction/latest?workflow_id=${workflowId}`,
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
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/workflow/activation`,
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
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
    const response = await fetch(`http://${apiEndpoint}:8080/api/workflow`, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      method: 'DELETE',
      body: JSON.stringify({
        workflow_id: workflowId,
        name: workflowName,
        action_id: actionId,
        reaction_id: reactionId,
      }),
    });
    if (response.status !== 200) console.error('Error invalide token');
    return true;
  } catch (error) {
    console.error('Error put Workflows data:', error);
    return false;
  }
}

export async function putWorkflows(
  apiEndpoint: string,
  token: string,
  formData: {
    workflow_id: number;
    name: string;
    action_option: any;
    reaction_option: any;
  },
) {
  try {
    const response = await fetch(
      `http://${apiEndpoint}:8080/api/user/workflows`,
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        method: 'PUT',
        body: JSON.stringify({
          workflow_id: formData.workflow_id,
          name: formData.name,
          action_option: formData.action_option,
          reaction_option: formData.reaction_option,
        }),
      },
    );
    if (response.status !== 200) {
      // console.error('Error invalide token');
      return false;
    }
    return true;
  } catch (error) {
    console.error('Error fetching Workflows data:', error);
    return false;
  }
}
