import React from 'react';

export default function DashboardPage() {
  return (
    <div className="flex flex-col gap-6 p-8 bg-slate-50 min-h-screen">
      <header className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-slate-900">Operations Dashboard</h1>
          <p className="text-slate-500">Real-time platform overview for TelcoFlow</p>
        </div>
        <div className="flex gap-3">
          <button className="px-4 py-2 bg-blue-600 text-white rounded-md shadow hover:bg-blue-700 transition">
            Create Order
          </button>
        </div>
      </header>

      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <StatCard title="Active Subscribers" value="1,245,678" change="+12.5%" />
        <StatCard title="Active Alarms" value="42" change="-5%" trend="down" />
        <StatCard title="Pending Orders" value="856" change="+3%" />
        <StatCard title="Revenue (MTD)" value="$12.4M" change="+8.2%" />
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="p-6 bg-white rounded-xl shadow-sm border border-slate-200 h-80">
          <h3 className="text-lg font-semibold mb-4 text-slate-800">Order Throughput</h3>
          <div className="flex items-center justify-center h-full text-slate-400 italic">
            [Chart Component Placeholder: Order Volume over 24h]
          </div>
        </div>
        <div className="p-6 bg-white rounded-xl shadow-sm border border-slate-200 h-80">
          <h3 className="text-lg font-semibold mb-4 text-slate-800">Critical Alarms</h3>
          <div className="space-y-4">
            <AlarmRow severity="critical" title="Link Failure: OLT-NYC-01" time="2 mins ago" />
            <AlarmRow severity="critical" title="High CPU: Core-Router-02" time="15 mins ago" />
            <AlarmRow severity="major" title="BGP Flap: Partner-Peering" time="45 mins ago" />
          </div>
        </div>
      </div>
    </div>
  );
}

function StatCard({ title, value, change, trend = 'up' }: any) {
  return (
    <div className="p-6 bg-white rounded-xl shadow-sm border border-slate-200">
      <p className="text-sm font-medium text-slate-500">{title}</p>
      <div className="mt-2 flex items-baseline justify-between">
        <p className="text-2xl font-bold text-slate-900">{value}</p>
        <span className={\`text-xs font-semibold \${trend === 'up' ? 'text-green-600' : 'text-red-600'}\`}>
          {change}
        </span>
      </div>
    </div>
  );
}

function AlarmRow({ severity, title, time }: any) {
  const colors: any = {
    critical: 'bg-red-500',
    major: 'bg-orange-500',
    minor: 'bg-yellow-500'
  };
  return (
    <div className="flex items-center justify-between p-3 bg-slate-50 rounded-lg">
      <div className="flex items-center gap-3">
        <div className={\`w-2 h-2 rounded-full \${colors[severity]}\`}></div>
        <span className="text-sm font-medium text-slate-700">{title}</span>
      </div>
      <span className="text-xs text-slate-400">{time}</span>
    </div>
  );
}
