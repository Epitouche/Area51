import { Text, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { AppContext } from '../context/AppContext';
import { useContext, useEffect } from 'react';
import { ServiceCard } from '../components';
import { parseServices } from '../service';

export default function ServiceScreen() {
  const {
    isBlackTheme,
    servicesConnected,
    aboutJson,
    serverIp,
    setServicesConnected,
  } = useContext(AppContext);

  useEffect(() => {
    if (aboutJson) parseServices({ aboutJson, serverIp, setServicesConnected });
  }, [serverIp]);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
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
            servicesConnected.services.length > 0 &&
            servicesConnected.services.map((service, index) => (
              <ServiceCard
                key={index}
                title={service.name}
                image={
                  'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000'
                }
                status={service.isConnected}
                handleOauthLogin={() => console.log('pressed')}
              />
            ))}
        </View>
      </View>
    </View>
  );
}
