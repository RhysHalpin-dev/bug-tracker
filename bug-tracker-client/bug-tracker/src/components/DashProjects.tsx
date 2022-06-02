import React from 'react';
import styled from 'styled-components';

const DashProjects: React.FC = () => {
  return <Container>PROJECTS</Container>;
};

export default DashProjects;

const Container = styled.div`
  border: solid 3px blue;
  width: 90%;
  display: flex;
  flex-direction: row;
  align-items: stretch;
`;
