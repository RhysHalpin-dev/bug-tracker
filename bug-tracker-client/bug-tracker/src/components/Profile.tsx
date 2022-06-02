import styled from 'styled-components';
import { useState, useContext, useEffect } from 'react';
import { UserContext } from '../context/userContext';
interface ProfileDetails {
  name?: string;
  email?: string;
  img?: string;
  bio?: string;
  contact?: string;
  admin?: boolean;
}

const profiles = [
  {
    name: 'Rhys Halpin',
    img: '/profileTest.png',
    description:
      'I am a full stack engineer, Current project lead of Bug-Tracker and currently learning Golang and React',
    contact: 'rhyshalpin@hotmail.co.uk',
    admin: true,
  },
];

const Profile: React.FC<ProfileDetails> = () => {
  const { state, actions } = useContext(UserContext);
  console.log(state.user.client_id);

  const [data, setData] = useState<ProfileDetails>();

  useEffect(() => {
    fetchProfile();
  }, []);

  const fetchProfile = async () => {
    try {
      const res = await fetch('http://localhost:8000/apiv1/auth/profile', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          userObject: state.user.client_id,
        }),
      });
      const data = await res.json();
      console.log(data);
      setData(data);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <Container>
      <h2>Profile</h2>

      <img src={profiles[0].img} alt="profileImg" />
      <p>{data?.name}</p>
      <p>{data?.bio}</p>
      <p>Role: {profiles[0].admin === true ? 'Admin' : 'Engineer'}</p>
      <p>{data?.email}</p>
      <button>Edit Profile</button>
    </Container>
  );
};

export default Profile;

const Container = styled.div`
  border: 3px solid pink;
  width: 30%;
  word-wrap: break-word;
  height: 50%;

  img {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    border: solid 3px;
  }
`;
