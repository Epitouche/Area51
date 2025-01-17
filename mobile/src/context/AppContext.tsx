// src/context/AppContext.tsx
import React, { createContext, useState, ReactNode, useEffect } from 'react';
import { AboutJson, AboutJsonParse, Workflow } from '../types';
import { getAboutJson, checkToken, getToken, saveToken } from '../service';

interface AppContextProps {
  serverIp: string;
  setServerIp: (ip: string) => void;
  aboutJson: AboutJson | undefined;
  setAboutJson: (aboutjson: AboutJson) => void;
  isConnected: boolean;
  setIsConnected: (isConnected: boolean) => void;
  isBlackTheme: boolean;
  setIsBlackTheme: (isBlackTheme: boolean) => void;
  servicesConnected: AboutJsonParse;
  setServicesConnected: (servicesConnected: AboutJsonParse) => void;
  workflows: Workflow[];
  setWorkflows: (workflows: Workflow[]) => void;
}

const AppContext = createContext<AppContextProps>({
  serverIp: '',
  setServerIp: () => {},
  aboutJson: undefined,
  setAboutJson: () => {},
  isConnected: false,
  setIsConnected: () => {},
  isBlackTheme: true,
  setIsBlackTheme: () => {},
  servicesConnected: {
    services: [],
  },
  setServicesConnected: () => { },
  workflows: [],
  setWorkflows: () => { },
});

interface AppProviderProps {
  children: ReactNode;
}

export default function AppProvider({ children }: AppProviderProps) {
  const [serverIp, setServerIp] = useState<string>('');
  const [aboutJson, setAboutJson] = useState<AboutJson>();
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [isBlackTheme, setIsBlackTheme] = useState<boolean>(false);
  const [servicesConnected, setServicesConnected] = useState<AboutJsonParse>(
    {} as AboutJsonParse,
  );
  const [workflows, setWorkflows] = useState<Workflow[]>([]);

  useEffect(() => {
    const checkConnection = async () => {
      if (await checkToken('token')) setIsConnected(true);
    };
    checkConnection();
    const aboutJson = async () => {
      if (serverIp != '') {
        getAboutJson(serverIp, setAboutJson);
      }
    };
    aboutJson();
  }, [serverIp]);

  useEffect(() => {
    const checkAndGrapServerIp = async () => {
      if (await checkToken('serverIp')) getToken('serverIp', setServerIp);
    };
    checkAndGrapServerIp();
  }, []);

  return (
    <AppContext.Provider
      value={{
        serverIp,
        setServerIp,
        aboutJson,
        setAboutJson,
        isConnected,
        setIsConnected,
        isBlackTheme,
        setIsBlackTheme,
        servicesConnected,
        setServicesConnected,
        workflows,
        setWorkflows,
      }}>
      {children}
    </AppContext.Provider>
  );
}

export { AppContext };
