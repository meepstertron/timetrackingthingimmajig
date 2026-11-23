import './App.css'
import { Bar, BarChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts';
import homeIcon from './assets/icons/home.svg';
import userIcon from './assets/icons/user.svg';
import projectIcon from './assets/icons/project.svg';
import streakIcon from './assets/icons/streak.webp'

type TimeMetric = {
    name: string;
    short: string;
    minutes: number;
};

const timeMetrics: TimeMetric[] = [
    { name: 'Total Observed Time', short: 'Observed', minutes: 204 },         // 3h 24min
    { name: 'Autonomous Agent Work Time', short: 'Agent Work', minutes: 72 }, // 1h 12min
    { name: 'Counted Effort Time', short: 'Effort', minutes: 98 },            // 1h 38min
];

function formatMinutes(total: number): string {
    const hours = Math.floor(total / 60);
    const minutes = total % 60;
    return hours > 0 ? `${hours}h ${minutes}min` : `${minutes}min`;
}

function TimeTooltip({ active, payload }: { active?: boolean; payload?: ReadonlyArray<{ payload: TimeMetric }> }) {
    if (!active || !payload || payload.length === 0) return null;
    const metric = payload[0].payload;
    return (
        <div className="bg-card rounded-md outline-2 outline-edge px-3 py-1.5">
            <span className="text-xs text-muted">{metric.name}</span>
            <p className="text-sm">{formatMinutes(metric.minutes)}</p>
        </div>
    );
}

export default function App() {
  return (
      <div className="bg-surface w-screen h-screen flex flex-col text-white">
          <div className="w-full h-12 flex flex-row justify-end px-8 items-center gap-4">
              <div className=" flex items-center gap-2"><span>Daily goal:</span> <progress max={100} value={50} className="rounded-full h-2 bg-panel "></progress></div>
              <div className="  outline-2  outline-streak-strong rounded-full h-6 w-fit px-2 text-sm text-white font-semibold items-center flex">🔥 2d </div> {/*TODO: Replace with real flame icon*/}

          </div>
          <div className="flex flex-row h-full">
              <div className='w-64 h-full px-2 flex flex-col gap-2 justify-between '>
                  <div className="flex flex-col gap-2">
                      <button className="w-full hover:bg-card h-12 rounded-md flex flex-row items-center gap-3 px-4 cursor-pointer"><img src={homeIcon} width={32}></img><span>Home</span></button>
                      <button className="w-full hover:bg-card h-12 rounded-md flex flex-row items-center gap-3 px-4 cursor-pointer"><img src={projectIcon} width={32}></img><span>Projects</span></button>
                      <button className="w-full hover:bg-card h-12 rounded-md flex flex-row items-center gap-3 px-4 cursor-pointer"><img src={streakIcon} width={32}></img><span>Streak</span></button>
                      <button className="w-full hover:bg-card h-12 rounded-md flex flex-row items-center gap-3 px-4 cursor-pointer"><img src={userIcon} width={32}></img><span>Profile</span></button>

                  </div>

                  <div className="flex flex-row items-center gap-2 mb-4">
                      <div className="rounded-lg overflow-hidden w-12 h-12 shrink-0">
                          <img
                              src="https://ca.slack-edge.com/E09V59WQY1E-U087PR1B2HX-e591611483e0-48"
                              alt="Jan Koch"
                              className="w-full h-full object-cover"
                          />
                      </div>
                      <div className="flex flex-col flex-1 min-w-0">
                          <span>Jan Koch</span>
                          <span>@jan</span>
                      </div>
                  </div>
              </div>
              <main className="bg-panel w-full h-full rounded-tl-md outline-edge outline-2 py-4 px-6">
                  <span className="text-muted">Last Updated: 2:27PM - 21.11.2026</span>
                  <hr className="mb-4 mt-1 text-muted"/>
                  <div className="flex flex-row justify-between"><h2 className="">Hi Jan,</h2><button>Edit Dashboard</button></div>

                  <div className="flex flex-row gap-2 w-full mt-6 h-48"> {/*feature cards*/}
                      <div className="bg-card rounded-md p-4 mb-4 w-2xl h-full">
                          <span className="text-muted">Time Observed today</span>
                          <p className="text-2xl mb-4">3h 24min</p>

                          <span className="text-muted">Estimated Effort</span>
                          <p className="text-2xl">1h 53min</p>
                      </div>
                      <div className="bg-card rounded-md p-4 mb-4 w-full h-full flex flex-col justify-between">
                          <span>Time breakdown</span>
                          <div className="px-8 w-full h-8 flex flex-row gap-0.5 ">
                              <div className="w-full h-full bg-green-400 rounded-md"></div>
                              <div className=" h-full border-blue-400 bg-hatched border-2 rounded-md w-1/6"></div>
                              <div className=" h-full  border-2 rounded-md w-1/3"></div>
                          </div>

                          <div className="w-full h-fit mt-4 flex flex-row justify-evenly text-sm">
                              <div className="flex flex-row items-center h-fit">
                                  <div className="w-6 h-6 bg-green-400 mr-2 rounded-md "></div>
                                  <span>Time with effort</span></div>
                              <div className="flex flex-row items-center h-fit">
                                  <div className="w-6 h-6 border-blue-400 bg-hatched border-2 mr-2 rounded-md"></div>
                                  <span>Time AI worked alone</span></div>
                              <div className="flex flex-row items-center h-fit">
                                  <div className="w-6 h-6 border-2 mr-2 rounded-md"></div>
                                  <span>Uncounted work</span></div>
                          </div>
                      </div>
                      <div className="bg-card rounded-md p-4 mb-4 w-full h-full flex flex-col">

                      </div>
                  </div>

              </main>
          </div>
      </div>
  );
}