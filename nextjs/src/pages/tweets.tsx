import { NextPage } from "next";
import { Title } from "../components/Title";
import { Tweet as TweetModel } from "../utils/models";
import { Tweet } from "../components/Tweet";
import api from "../service/api";
import useSWR from "swr";

const fetcher = (url: string) => api.get(url).then((res) => res.data);

const TweetsPage: NextPage = () => {
  const { data: tweets } = useSWR<TweetModel[]>("tweets", fetcher, {
    refreshInterval: 5000,
  });

  return (
    <div>
      <Title>Tweets</Title>
      {tweets?.map((t, key) => (
        <Tweet key={key} tweet={t} />
      ))}
    </div>
  );
};

export default TweetsPage;
