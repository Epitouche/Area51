// src/context/AppContext.tsx
import React, { createContext, useState, ReactNode, useEffect } from 'react';
import { AboutJson, AboutJsonParse } from '../types';
import { checkToken, getToken, saveToken } from '../service';
import { getAboutJson } from '../service/getAboutJson';

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
  setServicesConnected: () => {},
});

interface AppProviderProps {
  children: ReactNode;
}

export default function AppProvider({ children }: AppProviderProps) {
  const [serverIp, setServerIp] = useState<string>('');
  const [aboutJson, setAboutJson] = useState<AboutJson>();
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [isBlackTheme, setIsBlackTheme] = useState<boolean>(true);
  const [servicesConnected, setServicesConnected] = useState<AboutJsonParse>(
    {} as AboutJsonParse,
  );

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
    saveToken('serverIp', serverIp);
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
      }}>
      {children}
    </AppContext.Provider>
  );
}

export { AppContext };
