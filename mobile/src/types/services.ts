import { AuthConfiguration } from "react-native-app-auth";

export interface SelectServicesParamsProps {
  serviceName: string;
  serverIp: string;
  sessionToken?: string;
}

export type AuthApiCall = {
  config: AuthConfiguration;
  setToken: (accessToken: string) => void;
};
