import { StyleSheet } from 'react-native';

export const globalStyles = StyleSheet.create({
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
  container: {
    width: '100%',
    padding: '5%',
    paddingTop: '20%',
    alignItems: 'center',
    gap: 20,
  },
  buttonColor: {
    backgroundColor: '#B454FD',
  },
  line: {
    width: '90%',
    height: 2,
    borderRadius: 2,
    marginBottom: 16,
  },

  // White Theme
  wallpaper: {
    flex: 1,
    alignItems: 'center',
    backgroundColor: '#E8E9E9',
  },
  text: {
    color: '#1A1A1A',
    fontSize: 16,
    fontWeight: 'bold',
  },
  title: {
    fontSize: 30,
    color: '#1A1A1A',
    fontWeight: 'bold',
  },
  subtitle: {
    fontSize: 20,
    color: '#1A1A1A',
    fontWeight: 'semibold',
  },
  input: {
    borderBottomWidth: 1,
    borderColor: '#1A1A1A',
    padding: 5,
    marginVertical: 10,
    fontSize: 16,
    color: 'black',
  },
  lineColor: {
    backgroundColor: '#1A1A1A',
  },

  // Black Theme
  wallpaperBlack: {
    flex: 1,
    alignItems: 'center',
    backgroundColor: '#1A1A1A',
  },
  textBlack: {
    color: '#1A1A1A',
    fontSize: 16,
    fontWeight: 'bold',
  },
  titleBlack: {
    fontSize: 30,
    color: '#F7FAFB',
    fontWeight: 'bold',
  },
  subtitleBlack: {
    fontSize: 20,
    color: '#F7FAFB',
    fontWeight: 'semibold',
  },
  inputBlack: {
    borderBottomWidth: 1,
    borderColor: '#F7FAFB',
    padding: 5,
    marginVertical: 10,
    fontSize: 16,
    color: 'white',
  },
  lineColorBlack: {
    backgroundColor: '#F7FAFB',
  },
});
