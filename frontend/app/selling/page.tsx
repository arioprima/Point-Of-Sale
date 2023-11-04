import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Selling Page",
  description: "Selling Page Pos Mania Mantap",
  
  // other metadata
};

const SellingPage = () => {
  return (
    <>
      <div className="h-[100vh] w-full flex gap-2">
        <div className="w-3/4 bg-red-100 min-h-full">
          <h1>
            Selling Page
          </h1>
        </div>
        <div className="w-1/4 bg-meta-6 min-h-full"></div>

      </div>
    </>
  );
};

export default SellingPage;
