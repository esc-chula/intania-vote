import moment from "moment";

export const calculateTime = (startAt: Date, endAt: Date): string | null => {
  const now = new Date();

  if (now > startAt && now < endAt) {
    const timeDiff = endAt.getTime() - now.getTime();
    const hours = Math.floor(
      (timeDiff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60),
    );
    const minutes = Math.floor((timeDiff % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((timeDiff % (1000 * 60)) / 1000);

    if (timeDiff > 0) {
      if (timeDiff < 24 * 60 * 60 * 1000) {
        return `${String(hours).padStart(2, "0")}:${String(minutes).padStart(
          2,
          "0",
        )}:${String(seconds).padStart(2, "0")}`;
      } else {
        const days = Math.floor(timeDiff / (1000 * 60 * 60 * 24));
        return `เหลือเวลาอีก ${days} วัน`;
      }
    }
  } else if (now < startAt) {
    const timeDiff = startAt.getTime() - now.getTime();
    const hours = Math.floor(
      (timeDiff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60),
    );
    const minutes = Math.floor((timeDiff % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((timeDiff % (1000 * 60)) / 1000);

    if (timeDiff > 0) {
      if (timeDiff < 24 * 60 * 60 * 1000) {
        return `เริ่มในอีก ${String(hours).padStart(2, "0")}:${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}`;
      } else {
        const days = Math.floor(timeDiff / (1000 * 60 * 60 * 24));
        if (days === 1) {
          return "เริ่มพรุ่งนี้";
        } else if (days > 1) {
          return `เริ่มในอีก ${days} วัน`;
        }
      }
    }
  } else if (now > endAt) {
    return moment(endAt).format("DD/MM/YYYY");
  }

  return null;
};

export const choicesGridClassName = (choicesLength: number) => {
  switch (choicesLength) {
    case 1:
      return "grid-cols-1";
    case 2:
      return "grid-cols-1 grid-rows-2 sm:grid-cols-2 sm:grid-rows-none";
    case 3:
      return "grid-cols-1 grid-rows-3 sm:grid-cols-2 sm:grid-rows-none";
    case 4:
      return "grid-cols-2 grid-rows-2 sm:grid-rows-none";
    case 5:
      return "grid-cols-1 grid-rows-5 sm:grid-cols-2 sm:grid-rows-none";
    default:
      break;
  }
};
