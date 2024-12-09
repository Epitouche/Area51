// src/context/AppContext.tsx
import React, { createContext, useState, ReactNode } from 'react';
import { AboutJson } from '../types';

interface AppContextProps {
  serverIp: string;
  setServerIp: (ip: string) => void;
  aboutjson: AboutJson | undefined;
  setAboutJson: (aboutjson: AboutJson) => void;
}

const AppContext = createContext<AppContextProps>({
  serverIp: '',
  setServerIp: () => { },
  aboutjson: undefined,
  setAboutJson: () => { },
});

interface AppProviderProps {
  children: ReactNode;
}

export default function AppProvider({ children }: AppProviderProps) {
  const [serverIp, setServerIp] = useState<string>('');
  const [aboutjson, setAboutJson] = useState<AboutJson>();


  return (
    <AppContext.Provider value={{ serverIp, setServerIp, aboutjson, setAboutJson }}>
      {children}
    </AppContext.Provider>
  );
}

export { AppContext };
