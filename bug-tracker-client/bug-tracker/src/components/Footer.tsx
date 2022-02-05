import React from 'react';
import styled from 'styled-components';

const Footer: React.FC = () => {
  return <Container> © 2021 Rhys Halpin. All Rights Reserved.</Container>;
};

export default Footer;

const Container = styled.div`
  margin-top: 1rem;
  text-align: center;
  opacity: 0.5;
`;
