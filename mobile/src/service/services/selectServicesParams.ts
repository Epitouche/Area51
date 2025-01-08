import { githubLogin } from './github';
import { microsoftLogin } from './microsoft';
import { spotifyLogin } from './spotify';

interface SelectServicesParamsProps {
  serviceName: string;
  serverIp: string;
}

export async function selectServicesParams({
  serverIp,
  serviceName,
}: SelectServicesParamsProps) {
  switch (serviceName) {
    case 'spotify':
      return await spotifyLogin(serverIp);
    case 'github':
      return await githubLogin(serverIp);
    case 'microsoft':
      return await microsoftLogin(serverIp);
    default:
      return false;
  }
}
