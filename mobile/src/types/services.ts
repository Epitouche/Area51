import { AuthConfiguration } from "react-native-app-auth";
import { AboutJson, AboutJsonParse } from "./aboutJson";

export interface SelectServicesParamsProps {
  serviceName: string;
  serverIp: string;
  sessionToken?: string;
}

export type AuthApiCall = {
  config: AuthConfiguration;
  setToken: (accessToken: string) => void;
};

export interface RefreshServicesProps {
  serverIp: string;
  setAboutJson: (aboutJson: AboutJson) => void;
  setServicesConnected: (services: AboutJsonParse) => void;
  aboutJson: AboutJson | undefined;
}
