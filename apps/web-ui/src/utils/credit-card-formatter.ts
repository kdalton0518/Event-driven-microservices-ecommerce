export const creditCardFormatter = {
  format(creditCard: string) {
    if (creditCard.length === 0) {
      return creditCard;
    }
    return creditCard.replace(
      /(\d{4})\s(\d{4})\s(\d{4})\s(\d{4})/g,
      "$1 $2 $3 $4"
    );
  },
  unformat(creditCard: string) {
    return creditCard.replace(/\D/g, "");
  },
};
