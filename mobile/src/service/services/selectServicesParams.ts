import { githubLogin } from './github';
import { microsoftLogin } from './microsoft';
import { spotifyLogin } from './spotify';

interface SelectServicesParamsProps {
  serviceName: string;
  serverIp: string;
  sessionToken?: string;
}

export async function selectServicesParams({
  serverIp,
  serviceName,
  sessionToken,
}: SelectServicesParamsProps) {
  switch (serviceName) {
    case 'spotify':
      return await spotifyLogin(serverIp, sessionToken);
    case 'github':
      return await githubLogin(serverIp, sessionToken);
    case 'microsoft':
      return await microsoftLogin(serverIp, sessionToken);
    default:
      return false;
  }
}
