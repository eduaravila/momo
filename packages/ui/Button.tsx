
type Props =  React.ButtonHTMLAttributes<HTMLButtonElement> &{
  children: React.ReactNode;
}

export const Button = (props: Props) => {
  const { children,...rest } = props;
  return <button {...rest}>{children}</button>;
};
