import { Text, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { AppContext } from '../context/AppContext';
import { useContext, useEffect, useState } from 'react';
import { ServiceCard } from '../components';
import { getToken, parseServices } from '../service';

export default function ServiceScreen() {
  const {
    isBlackTheme,
    servicesConnected,
    aboutJson,
    serverIp,
    setServicesConnected,
  } = useContext(AppContext);
  const [token, setToken] = useState('');

  useEffect(() => {
    const getMyToken = async () => {
      await getToken('token', setToken);
    }
    getMyToken();
    if (aboutJson) parseServices({ aboutJson, serverIp, setServicesConnected });
  }, [serverIp]);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}
          accessibilityLabel="Service Screen">
          Service Screen
        </Text>
        <View
          style={{
            width: '100%',
            flexWrap: 'wrap',
            flexDirection: 'row',
            justifyContent: 'center',
            gap: 10,
          }}>
          {servicesConnected &&
            aboutJson &&
            servicesConnected.services &&
            servicesConnected.services.map((service, index) => (
              <ServiceCard
                key={index}
                title={service.name}
                image={service.image}
                status={service.isConnected}
                isMobile={isBlackTheme}
                aboutJson={aboutJson}
                serverIp={serverIp}
                setServicesConnected={setServicesConnected}
                token={token}
              />
            ))}
        </View>
      </View>
    </View>
  );
}
