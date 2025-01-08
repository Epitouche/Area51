import React, { useEffect, useState } from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { Button } from 'react-native-paper';
import { globalStyles } from '../styles/global_style';
import { getToken, getWorkflows } from '../service';
import { Workflow } from '../types';

type WorkflowTableProps = {
  detailsModalVisible: boolean;
  setDetailsModalVisible: (detailsModalVisible: boolean) => void;
  workflows?: Workflow[];
  isBlackTheme?: boolean;
};

export function WorkflowTable({
  setDetailsModalVisible,
  detailsModalVisible,
  workflows,
  isBlackTheme,
}: WorkflowTableProps) {

  return (
    <ScrollView horizontal>
      <View style={styles.container}>
        <View
          style={[
            styles.headerRow,
            isBlackTheme
              ? globalStyles.secondaryLight
              : globalStyles.secondaryDark,
          ]}>
          <Text
            style={[
              styles.headerCell,
              isBlackTheme ? globalStyles.text : globalStyles.textBlack,
            ]}>
            Name
          </Text>
          <Text
            style={[
              styles.headerCell,
              isBlackTheme ? globalStyles.text : globalStyles.textBlack,
            ]}>
            Action
          </Text>
          <Text
            style={[
              styles.headerCell,
              isBlackTheme ? globalStyles.text : globalStyles.textBlack,
            ]}>
            Reaction
          </Text>
          <Text
            style={[
              styles.headerCell,
              isBlackTheme ? globalStyles.text : globalStyles.textBlack,
            ]}>
            Is Active
          </Text>
          <Text
            style={[
              styles.headerCell,
              isBlackTheme ? globalStyles.text : globalStyles.textBlack,
            ]}>
            Details
          </Text>
        </View>
        {workflows &&
          workflows.map((workflow, index) => (
            <View
              key={index}
              style={[
                styles.row,
                isBlackTheme
                  ? globalStyles.secondaryLight
                  : globalStyles.secondaryDark,
              ]}>
              <View style={styles.cell}>
                <Text style={isBlackTheme ? styles.text : styles.textBlack}>
                  {workflow.name}
                </Text>
              </View>
              <View style={styles.cell}>
                <Text style={isBlackTheme ? styles.text : styles.textBlack}>
                  {workflow.action_name}
                </Text>
              </View>
              <View style={styles.cell}>
                <Text style={isBlackTheme ? styles.text : styles.textBlack}>
                  {workflow.reaction_name}
                </Text>
              </View>
              <View style={styles.cell}>
                <Text style={isBlackTheme ? styles.text : styles.textBlack}>
                  {workflow.is_active ? 'Yes' : 'No'}
                </Text>
              </View>
              <View style={styles.cell}>
                <Button
                  onPress={() => setDetailsModalVisible(!detailsModalVisible)}>
                  <Text style={isBlackTheme ? styles.text : styles.textBlack}>
                    ...
                  </Text>
                </Button>
              </View>
            </View>
          ))}
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    padding: 10,
    width: '100%',
    borderColor: '#ccc',
    borderRadius: 5,
  },
  headerRow: {
    alignContent: 'center',
    flexDirection: 'row',
    borderTopLeftRadius: 5,
    borderTopRightRadius: 5,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  headerCell: {
    width: 100,
    flex: 1,
    padding: 10,
    fontWeight: 'bold',
    textAlign: 'center',
    color: '#333',
  },
  row: {
    alignContent: 'center',
    flexDirection: 'row',
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
    justifyContent: 'center',
  },
  textBlack: {
    textAlign: 'center',
    color: '#F7FAFB',
    fontWeight: 'bold',
  },
  text: {
    textAlign: 'center',
    color: '#1A1A1A',
    fontWeight: 'bold',
  },
  cell: {
    justifyContent: 'center',
    alignItems: 'center',
    width: 100,
  },
  cellValidate: {
    backgroundColor: 'green',
    justifyContent: 'center',
    alignItems: 'center',
    width: 100,
  },
  cellUnvalide: {
    backgroundColor: 'red',
    justifyContent: 'center',
    alignItems: 'center',
    width: 100,
  },
  buttonText: {
    color: '#fff',
  },
});
