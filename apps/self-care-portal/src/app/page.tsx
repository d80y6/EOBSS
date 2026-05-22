import React from 'react';

export default function SelfCareHome() {
  return (
    <div className="max-w-4xl mx-auto p-8 space-y-8">
      <header className="flex justify-between items-center">
        <h1 className="text-3xl font-bold">Welcome back, John</h1>
        <div className="w-10 h-10 bg-slate-200 rounded-full"></div>
      </header>

      <section className="p-6 bg-blue-600 text-white rounded-2xl shadow-lg flex justify-between items-center">
        <div>
          <p className="text-blue-100 uppercase text-xs font-bold tracking-wider">Current Balance</p>
          <p className="text-4xl font-bold mt-1">$45.20</p>
        </div>
        <button className="px-6 py-3 bg-white text-blue-600 rounded-xl font-bold">Top Up</button>
      </section>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="p-6 bg-white border rounded-2xl shadow-sm">
          <h3 className="font-bold text-lg">Your Services</h3>
          <ul className="mt-4 space-y-3">
            <li className="flex justify-between items-center p-3 bg-slate-50 rounded-xl">
              <span>Fiber 500Mbps</span>
              <span className="text-green-600 text-sm font-bold">Active</span>
            </li>
            <li className="flex justify-between items-center p-3 bg-slate-50 rounded-xl">
              <span>Mobile Unlimited</span>
              <span className="text-green-600 text-sm font-bold">Active</span>
            </li>
          </ul>
        </div>

        <div className="p-6 bg-white border rounded-2xl shadow-sm">
          <h3 className="font-bold text-lg">Usage This Month</h3>
          <div className="mt-4 space-y-4">
            <div>
              <div className="flex justify-between text-sm mb-1">
                <span>Data</span>
                <span>45GB / 100GB</span>
              </div>
              <div className="w-full h-2 bg-slate-100 rounded-full">
                <div className="w-[45%] h-full bg-blue-500 rounded-full"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
