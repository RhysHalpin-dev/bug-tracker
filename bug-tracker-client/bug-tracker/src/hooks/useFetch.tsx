import React from 'react';

export const useFetch = (url: string, options: any) => {
  const [res, setRes] = React.useState(null);
  const [error, setError] = React.useState(null);

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await fetch(url, options);
        const json = await res.json();

        setRes(json);
      } catch (err: any) {
        setError(err);
      }
    };
    fetchData();
  }, []);
  return { res, error };
};
