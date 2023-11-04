"use client";
import { useState } from "react";

const SellingPage = () => {
  const tabs = [
    { id: 0, label: "Semua" },
    { id: 1, label: "Sembako" },
    { id: 2, label: "Snack" },
    { id: 3, label: "Minuman" },
    { id: 4, label: "Disabled", disabled: true },
  ];

  const [activeTab, setActiveTab] = useState(0);

  const handleTabClick = (tabId) => {
    setActiveTab(tabId);
  };

  return (
    <>
      <div className="h-[100vh] w-full flex gap-2">
        <div className="w-3/4 bg-red-100 min-h-full">
          <h1 className="text-lg font-semibold mb-2">Selling Page</h1>

          <div className="border-b border-gray-200 dark:border-gray-700">
            <ul className="flex flex-wrap -mb-px text-sm font-medium text-center text-gray-500 dark:text-gray-400">
              {tabs.map((tab) => (
                <li key={tab.id} className="mr-2">
                  <a
                    href="#"
                    onClick={() => handleTabClick(tab.id)}
                    className={`inline-flex items-center justify-center p-4 border-b-2 ${
                      tab.id === activeTab
                        ? "border-blue-600 text-blue-600 dark:text-blue-500 dark:border-blue-500"
                        : "border-transparent"
                    } rounded-t-lg ${
                      tab.disabled
                        ? "cursor-not-allowed text-gray-400 dark:text-gray-500"
                        : "hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
                    } group`}
                  > <svg className="w-4 h-4 mr-2 text-blue-600 dark:text-blue-500" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 18">
                  <path d="M6.143 0H1.857A1.857 1.857 0 0 0 0 1.857v4.286C0 7.169.831 8 1.857 8h4.286A1.857 1.857 0 0 0 8 6.143V1.857A1.857 1.857 0 0 0 6.143 0Zm10 0h-4.286A1.857 1.857 0 0 0 10 1.857v4.286C10 7.169 10.831 8 11.857 8h4.286A1.857 1.857 0 0 0 18 6.143V1.857A1.857 1.857 0 0 0 16.143 0Zm-10 10H1.857A1.857 1.857 0 0 0 0 11.857v4.286C0 17.169.831 18 1.857 18h4.286A1.857 1.857 0 0 0 8 16.143v-4.286A1.857 1.857 0 0 0 6.143 10Zm10 0h-4.286A1.857 1.857 0 0 0 10 11.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 18 16.143v-4.286A1.857 1.857 0 0 0 16.143 10Z"/>
              </svg>
                    {tab.label}
                  </a>
                </li>
              ))}
            </ul>
      {/* Render content based on the active tab */}
      {activeTab === 0 && <div>Profile Content</div>}
      {activeTab === 1 && <div>Dashboard Content</div>}
      {activeTab === 2 && <div>Settings Content</div>}
      {activeTab === 3 && <div>Contacts Content</div>}
      {activeTab === 4 && <div>Disabled Content</div>}
          </div>
        </div>
        <div className="w-1/4 bg-meta-6 min-h-full"></div>
      </div>
    </>
  );
};

export default SellingPage;
